use tonic::{transport::Server, Request, Response, Status};

use monitor_service::monitor_server::{Monitor, MonitorServer};
use monitor_service::{PingRequest, PingResponse};

use std::time::SystemTime;


pub mod monitor_service {
    tonic::include_proto!("monitor"); // The string specified here must match the proto package name
}

#[derive(Debug, Default)]
pub struct MyMonitor {}

#[tonic::async_trait]
impl Monitor for MyMonitor {
    async fn ping(&self, request: Request<PingRequest>) -> Result<Response<PingResponse>, Status> {
        println!("Got a request: {:?}", request);

        let ct = SystemTime::now().duration_since(SystemTime::UNIX_EPOCH).unwrap();
        let cs = ct.as_secs();
        let cn = ct.as_nanos();

        if cs > i64::MAX as u64 {
            panic!("time to big");
        }

        let ct = ::prost_types::Timestamp {
            seconds: cs as i64,
            nanos: cn as i32
        };

        let reply = monitor_service::PingResponse {
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
        .add_service(MonitorServer::new(monitor))
        .serve(addr)
        .await?;

    Ok(())
}
