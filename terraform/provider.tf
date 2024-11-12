provider "aws" {
  region = "ap-south-1"
}

terraform {
  backend "rds" {
    region = "ap-south-1"
    key = "./terraform.tfstate"    
  }
}