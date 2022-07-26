fn main() -> Result<(), Box<dyn std::error::Error>> {
   tonic_build::configure()
        .out_dir("src/protos/")
        .compile(&["../api/monitor.proto", "../api/rsi.proto"],
                 &["../api/"])?;
   Ok(())
}
