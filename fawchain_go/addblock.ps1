param(
    [double]$Data
)

$body = @{
    Data = $Data
} | ConvertTo-Json

Invoke-WebRequest -Uri http://localhost:8080/ -Method POST -ContentType "application/json" -Body $body
