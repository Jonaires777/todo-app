import Input from '@/components/Input/Input'
import Image from 'next/image'

export default function Login() {
  return (
    <form className="flex flex-col h-screen items-center justify-center p-2" style={{ background: 'linear-gradient(to right, #F8B205, #FF3366)'}}>
      <strong className='text-[#ffffff]'>Log in to your account!</strong>
      <div className='bg-[#F8B205] flex flex-col items-center justify-center gap-10 p-4 rounded-md shadow-md h-1/2 border-2'>
          <Input id='email' name='email' label='Email' type='email'/>
          <Input id='password' name='password' label='Password' type='password'/>
      </div>
    </form>
  )
}