extern crate protoc_rust;

fn main() {
    protoc_rust::Codegen::new()
        .out_dir("src/protos")
        .inputs(&["../api/monitor.proto"])
        .include("../api/")
        .protoc_path("../../protoc/bin/protoc")
        .run()
        .expect("Running protoc failed.");
}
