use std::time::SystemTime;
use uuid::Uuid;
use tonic::{Request, Response, Status};

pub mod protos;
use protos::{
    base::RequestInfo,
    monitor::{PingRequest, PingResponse, monitor_server::Monitor},
};

pub fn gen_prost_ts() -> ::prost_types::Timestamp {
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

pub fn gen_debug_data(uuid: Option<String>) -> RequestInfo {
    let uuid = match uuid {
        Some(s) => s,
        None => Uuid::new_v4().as_simple().to_string(),
    };

    RequestInfo { 
        ts: Some(gen_prost_ts()),
        id: uuid,
    }
}

impl RequestInfo {
    pub fn get_uuid(&self) -> String {
        let RequestInfo { ts: _, id } = self;
        return id.clone();
    }
}

#[derive(Debug, Default)]
pub struct MyMonitor;

#[tonic::async_trait]
impl Monitor for MyMonitor {
    async fn ping(&self, request: Request<PingRequest>) -> Result<Response<PingResponse>, Status> {
        println!("Got a request: {:?}", request);

        let reply = PingResponse {
            response_time: Some(gen_prost_ts()),
            version: String::from("0.0.1")
        };

        Ok(Response::new(reply))
    }
}
/*
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn debug_data_contains_string() {
        let RequestInfo { 

}
*/
