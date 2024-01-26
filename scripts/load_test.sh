#!/bin/bash

# URL e Headers
URL="http://localhost:8080/"
HEADER1="X-Tenant-ID: customer"
HEADER2="X-Customer-ID: 29302660000177"
HEADER3="X-Shard-ID: 1"

# Parâmetros do Teste
DURATION="10s"  # Duração do teste
THREADS=4       # Número de threads (ajuste conforme necessário)
CONNECTIONS=300  # Número de conexões simultâneas (ajuste conforme necessário)

# Executando o teste de carga com wrk
wrk -c $CONNECTIONS -d $DURATION -t $THREADS --latency -H "$HEADER1" -H "$HEADER2" -H "$HEADER3" $URL
