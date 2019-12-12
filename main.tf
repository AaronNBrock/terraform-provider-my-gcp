provider "my-gcp" {
    # keyfile = "test"
    # project_id = "test"
}

resource "my-gcp_bucket" "my-server" {
    bucket_name = "potato-aaron-n-brock2"
}
