import axios from 'axios';
import { LegacyCharacterEncoding } from 'crypto';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react' 
import Layout from '../../components/Layouts';
import { UserPage } from '../../components/UserPage';
import { PostProps, User } from '../../interface';

const UsersPage: React.FC = ({}) =>{
    const router = useRouter();
    const intId = typeof router.query.id === "string" ? parseInt(router.query.id) : -1;
    const [ user, setUser ] = useState<User | null>(null)
    useEffect(() => {
        axios.get(`/api/users/${intId}`, {
            withCredentials: true
          })
          .then(res => {
           setUser(res.data)
          }).catch(err => console.log(err))
    },[intId, router]) 
    
    return (
            <Layout title='post page'>
             {user ? (
                <UserPage user={user} me={false} />
             ):(
                <div>loading...</div>
             )}
            </Layout>
    );
}

export default UsersPage