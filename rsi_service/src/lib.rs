use std::time::SystemTime;
use uuid::Uuid;
use tonic::{Request, Response, Status};
use rust_decimal::Decimal;
use rust_decimal::prelude::*;
use rust_decimal_macros::dec;
use tokio_test;

pub mod protos;
use protos::{
    base::DebugData,
    rsi::{PriceData, RsiData, rsi_server},
    monitor::{PingRequest, PingResponse, monitor_server},
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

fn gen_debug_data(uuid: Option<String>) -> DebugData {
    let uuid = match uuid {
        Some(s) => s,
        None => Uuid::new_v4().as_simple().to_string(),
    };

    DebugData { 
        ts: Some(gen_prost_ts()),
        uuid,
    }
}

impl DebugData {
    fn get_uuid(&self) -> String {
        let DebugData { ts: _, uuid } = self;
        return uuid.clone();
    }
}

fn calc_rsi(pd: Vec<Decimal>) -> String {
    let mut up: Vec<Decimal> = Vec::with_capacity(14);
    let mut down: Vec<Decimal> = Vec::with_capacity(14);

    let _ = pd.windows(2).inspect(|arr| {
        if arr[1] > arr[0] {
            up.push(arr[1] - arr[0]);
            down.push(Decimal::ZERO);
        } else if arr[1] < arr[0] {
            up.push(Decimal::ZERO);
            down.push(arr[0] - arr[1]);
        } else {
            up.push(Decimal::ZERO);
            down.push(Decimal::ZERO);
        }
    }).collect::<Vec<_>>();

    let rs = calc_smma(up) / calc_smma(down);

    let rsi = Decimal::from(100) - (Decimal::from(100)/(Decimal::from(1)+rs));

    return rsi.to_string();
}

fn calc_smma(pd: Vec<Decimal>) -> Decimal {
    pd.iter().enumerate().fold(dec!(0), |accum, e| {
        (accum * Decimal::from(e.0) + e.1)/(Decimal::from(e.0 + 1))
    })
}

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
    async fn get_rsi(&self, request: Request<PriceData>) -> Result<Response<RsiData>, Status> {
        let PriceData { pd, debug } = request.into_inner();

        if pd.len() < 3 || pd.len() > 80 {
            return Err(Status::new(tonic::Code::InvalidArgument, "too many or too few price values"));
        }

        let pd: Vec<Decimal> = pd.iter().map(|x| { Decimal::from_str(x).unwrap() }).collect();


        let rsival = calc_rsi(pd);

        if let Some(debug) = debug {
            let DebugData { ts: _, uuid } = debug;
            let debug = gen_debug_data(Some(uuid));

            return Ok(Response::new(RsiData { rsival, debug: Some(debug) }));
        } else {
            let debug = gen_debug_data(None);
            return Ok(Response::new(RsiData { rsival, debug: Some(debug) }));
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_calc_rsi() {
        assert_eq!(calc_rsi(vec![
                            dec!(10),
                            dec!(11.3),
                            dec!(10.2),
                            dec!(11.5),
                            dec!(11.8),
                            dec!(10.9)
        ]),"59.183673469387755102040816327");

        assert_eq!(calc_rsi(vec![
                            dec!(3451.59),
                            dec!(3532.12),
                            dec!(3545.91),
                            dec!(3670.85),
                            dec!(3580.32),
                            dec!(3556.94),
                            dec!(3639.40),
                            dec!(3687.15)
        ]), "75.417583840476498769908066813")
    }

    #[test]
    fn get_rsi_too_few_values() {
        let psd = PriceData {
            pd: vec![ "3451.59".to_string(), "3532.12".to_string(), ],
            debug: Some(gen_debug_data(None)),
        };

        let r = MyRsi::default();

        if let Err(stat) = tokio_test::block_on(<MyRsi as rsi_server::Rsi>::get_rsi(&r, Request::new(psd))) {
            assert_eq!(stat.code(), tonic::Code::InvalidArgument);
        } else {
            panic!("get_rsi did not return any error as expected");
        }
    }

    #[test]
    fn get_rsi_zero_values() {
        let psd = PriceData {
            pd: Vec::new(),
            debug: Some(gen_debug_data(None)),
        };

        let r = MyRsi::default();

        if let Err(stat) = tokio_test::block_on(<MyRsi as rsi_server::Rsi>::get_rsi(&r, Request::new(psd))) {
            assert_eq!(stat.code(), tonic::Code::InvalidArgument);
        } else {
            panic!("get_rsi did not return any error as expected");
        }
    }

    #[test]
    fn get_rsi_too_many_values() {
        let mut pd = Vec::with_capacity(81);

        for _ in 0..81 {
            pd.push("3639.94".to_string());
        }

        let psd = PriceData {
            pd,
            debug: Some(gen_debug_data(None)),
        };

        let r = MyRsi::default();

        if let Err(stat) = tokio_test::block_on(<MyRsi as rsi_server::Rsi>::get_rsi(&r, Request::new(psd))) {
            assert_eq!(stat.code(), tonic::Code::InvalidArgument);
        } else {
            panic!("get_rsi did not return any error as expected");
        }
    }

    #[test]
    fn get_rsi_success_case() {
        let pd1 = vec![
            "3451.59".to_string(),
            "3532.12".to_string(),
            "3545.91".to_string(),
            "3670.85".to_string(),
            "3580.32".to_string(),
            "3556.94".to_string(),
            "3639.40".to_string(),
            "3687.15".to_string()
        ];

        let debug = gen_debug_data(None);
        let uuid_orig = debug.get_uuid();

        let psd = PriceData {
            pd: pd1,
            debug: Some(debug),
        };

        let r = MyRsi::default();

        match tokio_test::block_on(<MyRsi as rsi_server::Rsi>::get_rsi(&r, Request::new(psd))) {
            Ok(rsidat) => {
                let rsidat = rsidat.into_inner();
                if let RsiData { rsival, debug: Some(DebugData { ts: _, uuid }) } = rsidat {
                    assert_eq!(
                        rsival,
                        "75.417583840476498769908066813".to_string()
                    );

                    assert_eq!(
                        uuid,
                        uuid_orig
                    );
                } else {
                    panic!("get_rsi did not return debug data");
                }
            },
            Err(e) => {
                panic!("get_rsi returned error when not supposed to: {:?}", e);
            }
        }
    }
}
