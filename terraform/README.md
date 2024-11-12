# Terraform EKS, RDS, and Helm Project

## Overview

This project sets up an AWS EKS cluster with associated networking, deploys a MySQL database using RDS and Helm, and deploys a GoFiber backend application. It also includes NGINX as an ingress controller.

## Directory Structure

- `providers.tf`: Configures AWS, Kubernetes, and Helm providers.
- `backend.tf`: Configures Terraform backend using S3.
- `variables.tf`: Defines input variables.
- `outputs.tf`: Defines output values.
- `vpc.tf`: Sets up VPC, subnets, and networking.
- `security_groups.tf`: Defines security groups for EKS and RDS.
- `eks.tf`: Configures EKS cluster and node groups.
- `rds.tf`: Sets up RDS MySQL instance.
- `kubernetes.tf`: Deploys Kubernetes resources like GoFiber backend.
- `helm.tf`: Deploys Helm charts for NGINX and MySQL.
- `versions.tf`: Specifies Terraform and provider versions.
- `terraform.tfvars`: (Optional) Provides variable values.
- `.gitignore`: Excludes sensitive files from version control.

## Setup Instructions

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/terraform-eks-rds-helm.git
   cd terraform-eks-rds-helm
