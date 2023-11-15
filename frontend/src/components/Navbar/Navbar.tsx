import { GrHomeRounded } from "react-icons/gr";
import Link from "next/link";

export default function Navbar() {
    const style = { color: "white", fontSize: "1.5em" }
    
    return (
    <nav className="flex flex-row justify-between items-center bg-[#F8B205] p-2">
      <Link href="/">
        <GrHomeRounded style={style}/>
      </Link>
      <div className="flex flex-row gap-4">
        <Link className="text-[#ffffff] hover:scale-110 ease-in duration-100" href="/create">
          Create Todo
        </Link>
        <Link className="text-[#ffffff] hover:scale-110 ease-in duration-100" href="/login">
          Log In
        </Link>
        <Link className="text-[#ffffff] hover:scale-110 ease-in duration-100" href="/signup">
          Sign Up
        </Link>
        <button className="text-[#ffffff] hover:scale-110 ease-in duration-100">Log Out</button>
      </div>
    </nav>
  );
}
