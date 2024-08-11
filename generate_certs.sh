#!/bin/bash

CERTS_DIR="./certs"

echo "Cleaning up existing certificates..."
rm -rf $CERTS_DIR/*
mkdir -p $CERTS_DIR

CA_SUBJ="/C=US/ST=State/L=City/O=Organization/OU=Unit/CN=ltp-plotter-ca"

# Function to generate key, CSR, and sign certificate
generate_cert() {
    local NAME=$1
    local SUBJ=$2
    local ALT_NAME=$3

    echo "Generating private key for ${NAME}..."
    openssl genpkey -algorithm RSA -out $CERTS_DIR/${NAME}.key -pkeyopt rsa_keygen_bits:2048

    echo "Generating certificate signing request (CSR) for ${NAME}..."
    openssl req -new -key $CERTS_DIR/${NAME}.key -out $CERTS_DIR/${NAME}.csr -subj "$SUBJ"

    if [ -n "$ALT_NAME" ]; then
        echo "Signing certificate for ${NAME} with CA..."
        openssl x509 -req -in $CERTS_DIR/${NAME}.csr -CA $CERTS_DIR/ca.crt -CAkey $CERTS_DIR/ca.key -CAcreateserial -out $CERTS_DIR/${NAME}.crt -days 365 -sha256 -extfile <(printf "subjectAltName=DNS:${ALT_NAME}")
    else
        echo "Signing certificate for ${NAME} with CA..."
        openssl x509 -req -in $CERTS_DIR/${NAME}.csr -CA $CERTS_DIR/ca.crt -CAkey $CERTS_DIR/ca.key -CAcreateserial -out $CERTS_DIR/${NAME}.crt -days 365 -sha256
    fi
}

# Generate the Certificate Authority (CA)
echo "Generating CA private key and certificate..."
openssl genpkey -algorithm RSA -out $CERTS_DIR/ca.key -pkeyopt rsa_keygen_bits:2048
openssl req -x509 -new -nodes -key $CERTS_DIR/ca.key -sha256 -days 365 -out $CERTS_DIR/ca.crt -subj "$CA_SUBJ"

# Generate Server Certificates with correct SANs
generate_cert "expression_plot_server" "/C=US/ST=State/L=City/O=Organization/OU=Unit/CN=expr-plot-service" "expr-plot-service"
generate_cert "data_plot_server" "/C=US/ST=State/L=City/O=Organization/OU=Unit/CN=data-plot-service" "data-plot-service"
generate_cert "api_gateway_client" "/C=US/ST=State/L=City/O=Organization/OU=Unit/CN=api-gateway-client" ""

# Verify the certificates
echo "Verifying certificates..."
for SERVICE in expression_plot_server data_plot_server api_gateway_client; do
    if ! openssl verify -CAfile $CERTS_DIR/ca.crt $CERTS_DIR/${SERVICE}.crt; then
        RED='\033[0;31m'
        NC='\033[0m' # No Color
        printf "${RED}Error: Certificate verification failed for ${SERVICE}${NC}\n"
        exit 1
    fi
done

echo "Certificates generated and verified successfully."
