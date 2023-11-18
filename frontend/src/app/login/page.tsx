'use-client';


import Footer from '@/components/Footer/Footer'
import Input from '@/components/Input/Input'
import Navbar from '@/components/Navbar/Navbar'
import Image from 'next/image'

export default function Login() {
  const handleSubmit = () => {
    console.log("Login realizado")
  }
  
  return (
    <div>
      <Navbar />
      <form className="flex flex-col h-screen items-center justify-center p-2">
        <div className='flex flex-row bg-[#F8B205] border-4 rounded-lg'>
          <div className='flex flex-col items-center justify-center gap-10 p-4 rounded-md shadow-xl' style={{ background: 'linear-gradient(to right, #F8B205, #FF3366)'}}>
              <strong className='text-[#ffffff]'>Log in to your account!</strong>
              <Input id='email' name='email' label='Email' type='email'/>
              <Input id='password' name='password' label='Password' type='password'/>
              <button className='hover:scale-110 ease-in duration-100'><strong className='text-[#ffffff] bg-[#FF3E17] p-2 rounded-lg'>Log in</strong></button>
          </div>
          <Image src="https://img.freepik.com/premium-photo/brown-mug-with-face-that-says-i-m-monster_900101-71318.jpg?size=338&ext=jpg&ga=GA1.1.386372595.1697673600&semt=ais" alt='' width={256} height={256}/>
        </div>
      </form>
      <Footer />
    </div>
  )
}