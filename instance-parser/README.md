#RDS Instance spec parsing Tool 

<br>
dependent packages

- github.com/aws/aws-sdk-go

`go run /path/to/rds_instance_parser.go -args file.tf`
<br><br>Output as following for locals by Terraform<br>
Issue:
Value of db.r5.4xlarge memory is 192GiB via API query, but 128GiB is correct value.
Actually, set direct on code for that value temporary.
<br><br>
**sample.tf**
```hcl-terraform
locals {
  instance_types = {
    db.m4.4xlarge   = {cpu_cores = 16,  memory = 68719476736}
    db.r4.4xlarge   = {cpu_cores = 16,  memory = 130996502528}
    db.r3.xlarge    = {cpu_cores = 4,   memory = 32749125632}
    db.t3.2xlarge   = {cpu_cores = 8,   memory = 34359738368}
    db.t2.micro     = {cpu_cores = 1,   memory = 1073741824}
    db.t2.small     = {cpu_cores = 1,   memory = 2147483648}
    db.r5.12xlarge  = {cpu_cores = 48,  memory = 412316860416}
    db.m1.large     = {cpu_cores = 2,   memory = 8053063680}
    db.t3.medium    = {cpu_cores = 2,   memory = 4294967296}
    db.r3.4xlarge   = {cpu_cores = 16,  memory = 130996502528}
    db.m5.xlarge    = {cpu_cores = 4,   memory = 17179869184}
    db.m3.2xlarge   = {cpu_cores = 8,   memory = 32212254720}
    db.m4.xlarge    = {cpu_cores = 4,   memory = 17179869184}
    db.r5.2xlarge   = {cpu_cores = 8,   memory = 68719476736}
    db.r5.4xlarge   = {cpu_cores = 16,  memory = 137438953472}
    db.r4.16xlarge  = {cpu_cores = 64,  memory = 523986010112}
    db.t3.xlarge    = {cpu_cores = 4,   memory = 17179869184}
    db.x1e.32xlarge = {cpu_cores = 128, memory = 4191888080896}
    db.m2.xlarge    = {cpu_cores = 2,   memory = 18360985190}
    db.m2.2xlarge   = {cpu_cores = 4,   memory = 36721970381}
    db.r4.2xlarge   = {cpu_cores = 8,   memory = 65498251264}
    db.r3.large     = {cpu_cores = 2,   memory = 16374562816}
    db.r4.large     = {cpu_cores = 2,   memory = 16374562816}
    db.m1.xlarge    = {cpu_cores = 4,   memory = 16106127360}
    db.m4.large     = {cpu_cores = 2,   memory = 8589934592}
    db.m4.10xlarge  = {cpu_cores = 40,  memory = 171798691840}
    db.t3.micro     = {cpu_cores = 2,   memory = 1073741824}
    db.m3.medium    = {cpu_cores = 1,   memory = 4026531840}
    db.m3.xlarge    = {cpu_cores = 4,   memory = 16106127360}
    db.r5.large     = {cpu_cores = 2,   memory = 17179869184}
    db.m1.small     = {cpu_cores = 1,   memory = 1825361101}
    db.t2.medium    = {cpu_cores = 2,   memory = 4294967296}
    db.m3.large     = {cpu_cores = 2,   memory = 8053063680}
    db.x1e.2xlarge  = {cpu_cores = 8,   memory = 261993005056}
    db.t2.xlarge    = {cpu_cores = 4,   memory = 17179869184}
    db.r3.8xlarge   = {cpu_cores = 32,  memory = 261993005056}
    db.m4.16xlarge  = {cpu_cores = 64,  memory = 274877906944}
    db.t2.2xlarge   = {cpu_cores = 8,   memory = 34359738368}
    db.m5.large     = {cpu_cores = 2,   memory = 8589934592}
    db.m4.2xlarge   = {cpu_cores = 8,   memory = 34359738368}
    db.x1e.xlarge   = {cpu_cores = 4,   memory = 130996502528}
    db.r3.2xlarge   = {cpu_cores = 8,   memory = 65498251264}
    db.t1.micro     = {cpu_cores = 1,   memory = 658203738}
    db.r4.xlarge    = {cpu_cores = 4,   memory = 32749125632}
    db.r4.8xlarge   = {cpu_cores = 32,  memory = 261993005056}
    db.t3.small     = {cpu_cores = 2,   memory = 2147483648}
    db.t2.large     = {cpu_cores = 2,   memory = 8589934592}
    db.m1.medium    = {cpu_cores = 1,   memory = 4026531840}
    db.x1e.4xlarge  = {cpu_cores = 16,  memory = 523986010112}
    db.r5.24xlarge  = {cpu_cores = 96,  memory = 824633720832}
    db.t3.large     = {cpu_cores = 2,   memory = 8589934592}
    db.x1.32xlarge  = {cpu_cores = 128, memory = 2095944040448}
    db.m5.24xlarge  = {cpu_cores = 96,  memory = 412316860416}
    db.m2.4xlarge   = {cpu_cores = 8,   memory = 73443940762}
    db.m5.2xlarge   = {cpu_cores = 8,   memory = 34359738368}
    db.r5.xlarge    = {cpu_cores = 4,   memory = 34359738368}
    db.m5.4xlarge   = {cpu_cores = 16,  memory = 68719476736}
    db.x1e.8xlarge  = {cpu_cores = 32,  memory = 1047972020224}
    db.m5.12xlarge  = {cpu_cores = 48,  memory = 206158430208}
    db.x1.16xlarge  = {cpu_cores = 64,  memory = 1047972020224}
    db.x1e.16xlarge = {cpu_cores = 64,  memory = 2095944040448}
  }
}
```