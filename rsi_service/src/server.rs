use tonic::{transport::Server, Request, Response, Status};

mod protos;
use protos::monitor::{PingRequest,PingResponse,monitor_server};

use std::time::SystemTime;

#[derive(Debug, Default)]
pub struct MyMonitor {}

#[tonic::async_trait]
impl monitor_server::Monitor for MyMonitor {
    async fn ping(&self, request: Request<PingRequest>) -> Result<Response<PingResponse>, Status> {
        println!("Got a request: {:?}", request);

        let ct = SystemTime::now().duration_since(SystemTime::UNIX_EPOCH).unwrap();
        let cs = ct.as_secs();
        let cn = ct.as_nanos();

        if cs > i64::MAX as u64 {
            panic!("timevalue exceeds int64");
        }

        let ct = ::prost_types::Timestamp {
            seconds: cs as i64,
            nanos: cn as i32
        };

        let reply = PingResponse {
            response_time: Some(ct),
            version: String::from("0.0.1")
        };

        Ok(Response::new(reply))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse()?;
    let monitor = MyMonitor::default();

    Server::builder()
        .add_service(monitor_server::MonitorServer::new(monitor))
        .serve(addr)
        .await?;

    Ok(())
}
