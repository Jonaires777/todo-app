import Input from '@/components/Input/Input'
import Image from 'next/image'

export default function Signup() {
  return (
    <form className="flex flex-col h-screen items-center justify-center p-2" style={{ background: 'linear-gradient(to right, #F8B205, #FF3366)'}}>
      <div>Doens't have an account ? <strong className='text-[#ffffff]'>Signup here!</strong></div>
      <div className='bg-[#F8B205] flex flex-col items-center justify-between p-4 rounded-md shadow-md h-1/2 border-2'>
          <Input id='name' name='email' label='Name' type='text'/>
          <Input id='email' name='email' label='Email' type='email'/>
          <Input id='password' name='password' label='Password' type='password'/>
      </div>
    </form>
  )
}