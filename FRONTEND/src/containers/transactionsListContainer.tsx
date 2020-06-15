import React, { useEffect, useState } from "react"
import TransactionsService from "../services/transactionsService";
import TranactionListContent from "../contents/transactionListContent";
import { Transaction } from "../model/types";


export default function TransactionContainer() {  

  const getNewTransaxtion = ()=>{ 
    const transa: Transaction = {
      id: "",
      source: {
        alias: "",
        id: "",
        number: 0,
        owner: {
          id: "",
          name: ""
        }
      },
      amount: 0,
      dateTime: new Date(),
      destination: {
        alias: "",
        id: "",
        number: 0,
        owner: {
          id: "", 
          name : ""
        }
      }
    }
   return transa

  }

  const [transactions, setTransactions] = useState([])
  const [newtrans, setNewTrans] = useState(getNewTransaxtion())

  useEffect(() => {

    TransactionsService().add("1", "1", {}).then((data) => { 
      console.log(data)
    })

    TransactionsService().getAll("1", "1").then((transactions) => { 
      setTransactions(transactions)
    })
  }, [])


  const saveChange = (newTransaction: Transaction) => { 
    let copy = transactions
    //@ts-ignore
    copy.push(newTransaction)

    setTransactions(copy)
  }



  const props = {
    transactions: transactions,
    saveChange: saveChange
  }


  return TranactionListContent(props, newtrans)
}