fn main() -> Result<(), Box<dyn std::error::Error>> {
   tonic_build::configure()
        .build_client(true)
        .build_server(true)
        .protoc_arg("--experimental_allow_proto3_optional")
        .out_dir("./src/protos/")
        .compile(&[
            "../api/monitor.proto",
            "../api/rsi.proto",
            "../api/exa.proto"
        ], &["../api/"])?;

   Ok(())
}
