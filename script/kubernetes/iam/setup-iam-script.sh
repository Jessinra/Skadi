aws iam create-role \
  --role-name eksSkadiCNIRole \
  --assume-role-policy-document file://"cni-role-trust-policy.json"

aws iam attach-role-policy \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy \
  --role-name eksSkadiCNIRole

aws eks update-addon \
  --region ap-southeast-1 \
  --cluster-name skadi \
  --addon-name vpc-cni \
  --service-account-role-arn arn:aws:iam::442662599070:role/eksSkadiCNIRole

aws iam create-role \
  --role-name eksSkadiNodeRole \
  --assume-role-policy-document file://"node-role-trust-policy.json"

aws iam attach-role-policy \
  --policy-arn arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy \
  --role-name eksSkadiNodeRole
aws iam attach-role-policy \
  --policy-arn arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly \
  --role-name eksSkadiNodeRole