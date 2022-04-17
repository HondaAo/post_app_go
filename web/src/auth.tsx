import axios from 'axios';
import React, { createContext, useEffect, useState } from 'react' 
import { User } from './interface';

type AuthContextProps = {
    currentUser: User | null | undefined;
  };

export const AuthContext = createContext<AuthContextProps>({currentUser: undefined});

const AuthProvider: React.FC = ({children}) =>{
    const [ currentUser, setCurrentUser ] = useState<User | null | undefined>(undefined)
    useEffect(() => {
        axios.get(`/api/user/me`, { withCredentials: true})
        .then((res) => {
        if(res.data.message == 'not authed'){
          setCurrentUser(null)
        }else{
          setCurrentUser(res.data)
        }})
        .catch(err => console.log(err)) 
    },[currentUser])
    return (
      <AuthContext.Provider value={{currentUser}}>
          {children}
      </AuthContext.Provider>
    );
}

export default AuthProvider