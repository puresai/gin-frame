wrk.method = "POST"
wrk.body  = '{"accountid":"122323", "integral_count":120}'
wrk.headers["Content-Type"] = "application/json"
function request()
    return wrk.format('POST', nil, nil, body)
end
