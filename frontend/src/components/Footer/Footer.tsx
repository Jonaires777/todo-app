import Link from "next/link"

export default function Footer(){
    return(
        <footer className="flex flex-col items-center gap-2 bg-[#F8B205]">
            <Link className="text-[#ffffff]" href="https://github.com/Jonaires777">See my Github account clicking here!</Link>
            <Link className="text-[#ffffff]" href="https://www.linkedin.com/in/jonathan-aires-2258a1268/">See my Linkedin account clicking here!</Link>
        </footer>
    )
}