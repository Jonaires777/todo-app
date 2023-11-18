'use-client';

import Footer from '@/components/Footer/Footer'
import Input from '@/components/Input/Input'
import Navbar from '@/components/Navbar/Navbar'
import Image from 'next/image'

export default function Signup() {
  const handleSubmit = () => {
    console.log("Form enviado")
  }
  
  return (
    <div>
      <Navbar />
      <form className="flex flex-col h-screen items-center justify-center p-2">
        <div className='flex flex-row bg-[#F8B205] rounded-lg border-4'>
          <div className='flex flex-col items-center justify-between p-4 rounded-md shadow-md' style={{ background: 'linear-gradient(to right, #F8B205, #FF3366)'}}>
              <div>Doens't have an account ? <strong className='text-[#ffffff]'>Signup here!</strong></div>
              <Input id='name' name='email' label='Name' type='text'/>
              <Input id='email' name='email' label='Email' type='email'/>
              <Input id='password' name='password' label='Password' type='password'/>
              <button type='submit' className='hover:scale-110 ease-in duration-100'><strong className='text-[#ffffff] bg-[#FF3E17] p-2 rounded-lg'>Sign Up</strong></button>
          </div>
          <Image className='rounded-lg' src="https://media.istockphoto.com/id/1176056835/ru/%D0%B2%D0%B5%D0%BA%D1%82%D0%BE%D1%80%D0%BD%D0%B0%D1%8F/%D0%BA%D0%B0%D0%B2%D0%B0%D0%B9%D0%B8-%D0%BA%D0%BE%D1%84%D0%B5%D0%B9%D0%BD%D0%B0%D1%8F-%D1%87%D0%B0%D1%88%D0%BA%D0%B0-%D1%85%D0%B0%D1%80%D0%B0%D0%BA%D1%82%D0%B5%D1%80-%D0%B2-%D1%83%D0%BF%D1%80%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B8-%D0%B4%D0%B5%D0%B9%D1%81%D1%82%D0%B2%D0%B8%D0%B9.jpg?s=612x612&w=0&k=20&c=B5BL631hCsxrYi6nfKqaCEYFjiO1I9VoVYqTbp3_dk8=" alt='' width={256} height={256}/>
        </div>
      </form>
      <Footer />
    </div>
  )
}