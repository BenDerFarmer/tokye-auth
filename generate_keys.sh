openssl ecparam -genkey -name secp521r1 -noout -out private.pem
openssl ec -in private.pem -pubout -out public.pem
