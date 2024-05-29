use std::{env::var, str::FromStr};

pub fn call_prover(param: String, function: &str) -> Option<String> {
    let prover_rpc = var("SHADOW_PROVING_PROVER_RPC").expect("Cannot detect PROVER_RPC env var");

    let client = reqwest::blocking::Client::new();
    let url = prover_rpc.to_owned() + function;
    let response = client
        .post(url)
        .header(
            reqwest::header::CONTENT_TYPE,
            reqwest::header::HeaderValue::from_static("application/json"),
        )
        .body(param.clone())
        .send();
    let rt: Result<String, reqwest::Error> = match response {
        Ok(x) => x.text(),
        Err(e) => {
            log::error!("call prover error, param =  {:#?}, error = {:#?}", param, e);
            return None;
        }
    };

    let rt_text = match rt {
        Ok(x) => x,
        Err(e) => {
            log::error!("fetch prover res_txt error, param =  {:#?}, error = {:#?}", param, e);
            return None;
        }
    };

    Some(rt_text)
}

pub fn read_env_var<T: Clone + FromStr>(var_name: &'static str, default: T) -> T {
    std::env::var(var_name)
        .map(|s| s.parse::<T>().unwrap_or_else(|_| default.clone()))
        .unwrap_or(default)
}

#[tokio::test]
async fn test_call_prover() {
    env_logger::Builder::from_env(env_logger::Env::default().default_filter_or("info")).init();
    dotenv::dotenv().ok();

    use crate::shadow_prove::ProveRequest;
    let request = ProveRequest {
        batch_index: 12,
        chunks: vec![vec![1], vec![2]],
        rpc: "http://localhost:3030".to_string(),
        shadow: false,
    };

    let rt = tokio::task::spawn_blocking(move || call_prover("1".to_string(), "/query_proof"))
        .await
        .unwrap();

    match rt {
        Some(info) => {
            if info.eq("success") {
                log::info!("successfully submitted prove task, waiting for proof to be generated");
            } else {
                log::error!("submitt prove task failed: {:#?}", info);
            }
        }
        None => {
            log::error!("submitt prove task failed");
        }
    }
}
