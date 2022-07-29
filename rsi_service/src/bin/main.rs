use tonic::{transport::Server, Request, Response, Status};

use rsi_service::protos::monitor::{PingRequest,PingResponse,monitor_server};
use rsi_service::protos::rsi::{PeriodLength,RsiData,rsi_server};
use rsi_service::protos::base::{DebugData};

use rsi_service::{gen_debug_data, gen_prost_ts};

#[derive(Debug, Default)]
pub struct MyMonitor {}

#[tonic::async_trait]
impl monitor_server::Monitor for MyMonitor {
    async fn ping(&self, request: Request<PingRequest>) -> Result<Response<PingResponse>, Status> {
        println!("Got a request: {:?}", request);

        let reply = PingResponse {
            response_time: Some(gen_prost_ts()),
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
        let PeriodLength { pl: _, debug } = request.into_inner();
        if let Some(debug) = debug {
            let DebugData { ts: _, uuid } = debug;
            let debug = gen_debug_data(Some(uuid))?;

            return Ok(Response::new(RsiData { rsival: 40, debug: Some(debug) }));
        } else {
            let debug = gen_debug_data(None)?;
            return Ok(Response::new(RsiData { rsival: 40, debug: Some(debug) }));
        }
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
