export function GetHeadersProperties(method:string, body?: any) : any {
    let headerSettings = {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Methods': 'DELETE, POST, GET, PUT, OPTIONS',
        'Access-Control-Allow-Headers': 'Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With'
    }
    const props: { method: string, headers: any, body?: any } ={ method: method, headers: headerSettings };

    if (body)
        props.body = JSON.stringify(body);
    
    return props
}
