import Head from 'next/head';
import  {Header}  from './Header';

interface LayoutsProps {
    title: string
}

const Layout: React.FC<LayoutsProps> = ({title, children}) =>{
        return (
        <>
        <Head>
        <title>{title}</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/favicon.ico" />
        </Head>
        <Header />
            <div className='w-4/6 m-4 mx-auto'>
                {children}
            </div>
        </>
        );
}
export default Layout