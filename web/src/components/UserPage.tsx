import React from 'react' 
import { User } from '../interface';

interface UserPageProps {
    user: User
    me: boolean
}

export const UserPage: React.FC<UserPageProps> = ({user, me}) =>{
    return (
       <div>
        <div className='my-4 font-bold'>
         {user.username}
        </div>
        <div className='my-4' id='bio'>
         {user.bio}
        </div>
       </div>
    );
}