use std::time::SystemTime;
use uuid::Uuid;
use tonic::Status;
use rust_decimal::{Decimal};
use rust_decimal_macros::dec;

pub mod protos;
use protos::base::DebugData;

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

pub fn calc_rsi(pd: Vec<Decimal>) -> String {
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }

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
}
