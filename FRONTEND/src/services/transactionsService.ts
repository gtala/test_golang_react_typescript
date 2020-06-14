import { GetHeadersProperties } from "../utils/getHeadersProperties"

export default function TransactionsService() {
  let host = "http://127.0.0.1:5000"  

  const add = (userId: string, accountId: string, body: any) => {
    const url = host + "/api/v1/user/" + userId + "/account/" + accountId + "/transactions"
    return fetch(url, GetHeadersProperties("POST", body)).then(r => (!r.ok && r.status !== 200) ? Promise.reject({ message: "An error occurred. " + r.status }) : r.json())
  }

    return {
      add: add
    }
}