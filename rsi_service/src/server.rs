use uuid::Uuid;
use tonic::{transport::Server, Request, Response, Status};

mod protos;
use protos::monitor::{PingRequest,PingResponse,monitor_server};
use protos::rsi::{PeriodLength,RsiData,rsi_server};
use protos::base::DebugData;

use std::time::SystemTime;

fn get_prost_ts() -> ::prost_types::Timestamp {
    let ct = SystemTime::now().duration_since(SystemTime::UNIX_EPOCH).unwrap();
    let mut cs = ct.as_secs();
    let cn = ct.as_nanos() - ((ct.as_secs() * 1000000000) as u128);

    // This ping function will serve inaccurate timestamps starting from Fri, 11 Apr 2262 23:47:16 UTC
    if cs > i64::MAX as u64 {
        cs = i64::MAX as u64;
    }

    ::prost_types::Timestamp {
        seconds: cs as i64,
        nanos: cn as i32
    }
}


#[derive(Debug, Default)]
pub struct MyMonitor {}

#[tonic::async_trait]
impl monitor_server::Monitor for MyMonitor {
    async fn ping(&self, request: Request<PingRequest>) -> Result<Response<PingResponse>, Status> {
        println!("Got a request: {:?}", request);

        let reply = PingResponse {
            response_time: Some(get_prost_ts()),
            version: String::from("0.0.1")
        };

        Ok(Response::new(reply))
    }
}

#[derive(Debug, Default)]
pub struct MyRsi {}

#[tonic::async_trait]
impl rsi_server::Rsi for MyRsi {
    async fn get_rsi(&self, request: Request<PeriodLength>) -> Result<Response<RsiData>, Status> {
        let uuid = match std::str::from_utf8(Uuid::new_v4().as_bytes()) {
            Ok(s) => String::from(s),
            Err(e) => return Err(Status::new(::tonic::Code::Aborted, 
                format!("generated UUID contains invalid utf8: {}", e))),
        };

        let debug = DebugData { 
            ts: Some(get_prost_ts()),
            uuid: String::from(uuid),
        };
        Ok(Response::new(RsiData { rsival: 40, debug: Some(debug) }))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse()?;
    let monitor = MyMonitor::default();
    let rsi = MyRsi::default();

    Server::builder()
        .add_service(monitor_server::MonitorServer::new(monitor))
        .add_service(rsi_server::RsiServer::new(rsi))
        .serve(addr)
        .await?;

    Ok(())
}
