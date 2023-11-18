"use-client"

import Footer from '@/components/Footer/Footer'
import Input from '@/components/Input/Input'
import Navbar from '@/components/Navbar/Navbar'
import Image from 'next/image'

export default function Create() {
  const handleSubmit = () => {
    console.log("Todo criado")
  }  
  
  return (
    <div>
      <Navbar />
      <form className="flex flex-col h-screen items-center justify-center p-2">
        <div className='flex flex-row shadow-xl bg-[#F8B205] border-4 rounded-lg'>
          <div className='flex flex-col items-center justify-center gap-10 p-4 rounded-md shadow-md' style={{ background: 'linear-gradient(to right, #F8B205, #FF3366)'}}>
              <strong className='text-[#ffffff]'>Create your todos here!</strong>
              <Input id='title' name='title' label='Title' type='email'/>
              <textarea className='p-2' placeholder='Content...'></textarea>
              <button className='hover:scale-110 ease-in duration-100'><strong className='text-[#ffffff] bg-[#FF3E17] p-2 rounded-lg'>Create Todo</strong></button>
          </div>
          <Image className='rounded-lg' src="https://img.freepik.com/free-vector/hand-drawn-coffee-cartoon-illustration_23-2150631722.jpg?size=338&ext=jpg&ga=GA1.1.1826414947.1699833600&semt=ais" alt='Todos' width={256} height={256}/>
        </div>
      </form>
      <Footer />
    </div>
  )
}