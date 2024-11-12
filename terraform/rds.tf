resource "aws_db_instance" "Fiber Feed RDS Instance" {
    allocated_storage = 20
    engine = "mysql"
    engine_version = "8.0"
    instance_class = "db.t3.micro"
    name = "Fiber Feed RDS"
    username = var.db_username
    password = var.db_password
    parameter_group_name = "default.mysql8.0"
    skip_final_snapshot = true 
    publicly_accessible = false
    vpc_security_group_ids = [aws_security_group.rds_sg.id]
    subnet_ids = var.db_subnet_ids
}