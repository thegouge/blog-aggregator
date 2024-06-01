'use client'

export default function RegisterForm() {

	async function registerUser(formData: FormData) {
		const response = await fetch("http://localhost:8000/v1/users", {
			method: "POST",
			body: JSON.stringify({
				name: formData.get("userName"),
				password: formData.get("Password")
			}),
		})
		console.log({ response })
	}

	return (
		<form action={registerUser} className="flex-column">
			<label For="userName" className="block">
				User Name:
				<input type="text" name="userName" className="block" />
			</label>
			<label For="password" className="block">
				Password:<input type="password" name="password" className="block" />
			</label>
			<input type="submit" value="Submit!" />
		</ form>
	)
}
