use tonic::transport::Server;
use utils::{
    MyMonitor,
    protos::{
        monitor::monitor_server::MonitorServer,
        rsi::rsi_server::RsiServer,
    }
};
use rsi_service::MyRsi;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse()?;
    let monitor = MyMonitor::default();
    let rsi = MyRsi::default();

    Server::builder()
        .add_service(MonitorServer::new(monitor))
        .add_service(RsiServer::new(rsi))
        .serve(addr)
        .await?;

    Ok(())
}
