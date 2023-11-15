import Footer from "@/components/Footer/Footer";
import Navbar from "@/components/Navbar/Navbar";
import Image from "next/image";

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col justify-between">
      <Navbar />
      <div className="flex flex-col gap-2 items-center">
        <Image
          className="rounded-lg"
          src="https://media.istockphoto.com/id/1175851148/vector/kawaii-coffee-cup-character-in-meditate-pose.jpg?s=612x612&w=0&k=20&c=6xn41fxI3T2_TIsRzZg0b6fmc2riNLN3JpS_7802ba8="
          alt="Relax"
          width={256}
          height={256}
        />
        <h1>
          <strong>Enjoy your day, Marcelito</strong>
        </h1>
        <h2>
          Today you have completed all your tasks and have achieved the{" "}
          <strong>#StateZer0!</strong>
        </h2>
      </div>
      <Footer />
    </main>
  );
}
