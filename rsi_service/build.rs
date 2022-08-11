fn main() -> Result<(), Box<dyn std::error::Error>> {
   tonic_build::configure()
        .build_client(false)
        .out_dir("./src/protos/")
        .compile(&["../api/monitor.proto", "../api/rsi.proto"],
                 &["../api/"])?;
   Ok(())
}
