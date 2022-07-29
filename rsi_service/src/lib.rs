use uuid::Uuid;
use tonic::Status;
pub mod protos;
use protos::base::DebugData;
use std::time::SystemTime;

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

pub fn gen_debug_data(uuid: Option<String>) -> Result<DebugData, Status> {
    let uuid = match uuid {
        Some(s) => s,
        None => 
            match std::str::from_utf8(Uuid::new_v4().as_bytes()) {
                Ok(s) => String::from(s),
                Err(e) => 
                    return Err(Status::new(::tonic::Code::Aborted, 
                        format!("generated UUID contains invalid utf8: {}", e))),
            }
    };

    Ok(DebugData { 
        ts: Some(gen_prost_ts()),
        uuid,
    })
}
#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}
