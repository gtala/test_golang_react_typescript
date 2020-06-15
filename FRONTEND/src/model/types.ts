
export interface User {
  id: string,
  name: string
}

export interface Account { 
  id: string,
  alias: string,
  number: number,
  owner: User
}

export interface Transaction { 
  id: string,
  amount: number,
  source: Account,
  destination: Account,
  dateTime: Date
}