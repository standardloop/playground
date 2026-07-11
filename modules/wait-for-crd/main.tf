
resource "null_resource" "wait_for_crds" {
  for_each = toset(var.crds)
  provisioner "local-exec" {
    command = <<-EOT
        attempts=${var.attempts}
        while !kubectl wait --for condition=established --timeout=${var.timeout} crd ${each.value} > /dev/null 2>&1
        do
          sleep 5
          let attempts=$attempts-1
          if [ $attempts -eq 0 ]
          then
            exit 1
          fi
        done
    EOT
  }
}
