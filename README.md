version Go v3
Librerias utilizadas
- fiber
- gonum

Instrucciones para correr en desarrollo

comando "go run ."

-----
ruta de servicio google publicado

postman request POST 'https://go-service-953268766315.northamerica-northeast1.run.app/refactor-qr' \
  --header 'Content-Type: application/json' \
  --body '{
  "matrix": [
    [2,1],
    [1,-3],
    [5,1]
  ]
}'