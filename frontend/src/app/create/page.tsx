"use-client"

import Input from '@/components/Input/Input'
import Image from 'next/image'

export default function Create() {
    return (
    <form action="/" className="flex flex-col h-screen items-center justify-center p-2" style={{ background: 'linear-gradient(to right, #F8B205, #FF3366)'}}>
      <strong className='text-[#ffffff]'>Create your todos here!</strong>
      <div className='bg-[#F8B205] flex flex-col items-center justify-center gap-10 p-4 rounded-md shadow-md h-1/2 border-2'>
          <Input id='title' name='title' label='Title' type='email'/>
          <textarea className='p-2' placeholder='Content...'></textarea>
          <button className='hover:scale-110 ease-in duration-100'><strong className='text-[#ffffff] bg-[#FF3E17] p-2 rounded-lg'>Create Todo</strong></button>
      </div>
    </form>
  )
}