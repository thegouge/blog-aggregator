import Link from "next/link"

export default function Header() {
	return (
		<div className="z-10 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
			<h1>Blog Aggregator</h1>
			<div className="lg:flex justify-between">
				<Link href="/login">
					Log In!
				</Link>
				<div className="mx-2"></div>
				<Link href="/register">
					Register!
				</Link>

			</div>
		</div >
	)
}
