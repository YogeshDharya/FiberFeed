provider "kubernetes" {
  host = aws_eks_cluster.
}
#TODO above provider definition is incomplete 
resource "kubernetes_secret" "db_credentials" {
  metadata {
    name      = "db-credentials"
    namespace = "default"
  }

  data = {
    username = var.db_username
    password = var.db_password
  }
}

resource "kubernetes_deployment" "go_fiber_backend" {
  metadata {
    name = "Fiber Feed Deployment Metadata"
  }

  spec {
    replicas = 2

    template {
      metadata {
        labels = {
          app = "Fiber Feed Backend"
        }
      }

      spec {
        container {
          name  = "Fiber Feed Backend Container"
          image = "Fiber Feed Image  "

          env {
            name  = "DB_HOST"
            value = aws_db_instance.my_rds_instance.endpoint
          }

          env {
            name  = "DB_PORT"
            value = "3306"
          }

          env {
            name = "DB_USER"
            value_from {
              secret_key_ref {
                name = "db-credentials"
                key  = "username"
              }
            }
          }

          env {
            name = "DB_PASSWORD"
            value_from {
              secret_key_ref {
                name = "db-credentials"
                key  = "password"
              }
            }
          }
        }
      }
    }
  }
}
