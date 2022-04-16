import axios from 'axios';
import React, { useEffect, useState } from 'react' 
import Layout from '../components/Layouts';
import { UserPage } from '../components/UserPage';
import { User } from '../interface';

interface mypageProps {

}

const Mypage: React.FC<mypageProps> = ({}) =>{
    const [ user, setUser ] = useState<User | null>(null)
    useEffect(() => {
        axios.get(`/api/user/me`, { withCredentials: true })
        .then((res) => setUser(res.data))
        .catch((err) => console.log(err))
    },[])
    return (
        <Layout title='mypage'>
        {user ? (
            <UserPage user={user} me={true} />
        ):(
            <div>
                loading....
            </div>
        )}  
        </Layout>
    );
}

export default Mypage 