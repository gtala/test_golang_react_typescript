import React, { useEffect } from "react"
import TransactionsService from "../services/transactionsService";

export default function TransactionContainer() {  




  useEffect(() => {

    TransactionsService().add("1", "1", {}).then((data) => { 
      console.log(data)
    })
  }, [])



  return <h2>HOLA MON</h2>;
  



}