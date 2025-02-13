import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
	Card,
	CardContent,
	CardDescription,
	CardHeader,
	CardTitle,
	CardFooter,
} from "@/components/ui/card";
import { MinusCircle, PlusCircle } from "lucide-react";
import { BsPaypal, BsStripe } from "react-icons/bs";

export default function MysteryBoxProduct() {
	const [quantity, setQuantity] = useState(1);

	const handleQuantityChange = (e: React.ChangeEvent<HTMLInputElement>) => {
		const value = Number.parseInt(e.target.value);
		setQuantity(Number.isNaN(value) || value < 1 ? 1 : value);
	};

	const handlePayment = async (
		gateway: "stripe" | "paypal" | "mercadopago",
	) => {
		// This is where you'd implement the actual payment logic
		console.log(
			`Processing payment with ${gateway} for ${quantity} mystery box(es)`,
		);
		// You'd typically make an API call here to your server to initiate the payment
	};

	return (
		<Card className="w-full max-w-md mx-auto">
			<CardHeader>
				<CardTitle className="text-2xl">Mystery Box</CardTitle>
				<CardDescription>Unbox the unknown for just $0.01!</CardDescription>
			</CardHeader>
			<CardContent>
				<div className="flex justify-center mb-4">
					<img
						className="bg-white w-[200px] h-[200px]"
						src="./mystery-box.png"
						alt="mystery box"
					/>
				</div>
				<p className="text-sm text-gray-600 mb-4">
					What's inside? It could be anything! From digital goodies to discount
					codes, each mystery box is a surprise waiting to be unveiled.
				</p>
				<div className="flex flex-col items-center">
					<div className="flex justify-between w-full">
						<label htmlFor="quantity" className="font-medium">
							<span>Quantity:</span>
						</label>
						<div className="flex gap-2">
							<Button
								variant={"outline"}
								className="cursor-pointer"
								type="button"
								onClick={() => setQuantity((prev) => (prev > 1 ? prev - 1 : 1))}
								disabled={quantity === 1}
							>
								<MinusCircle />
							</Button>
							<Input
								type="tel"
								id="quantity"
								value={quantity}
								onChange={handleQuantityChange}
								min="1"
								className="w-15 text-center"
							/>
							<Button
								variant={"outline"}
								type="button"
								className="cursor-pointer"
								onClick={() => setQuantity((prev) => prev + 1)}
							>
								<PlusCircle />
							</Button>
						</div>
					</div>
					<p className="mt-4 text-lg font-medium">
						Total Value ${(quantity * 0.01).toFixed(2)}
					</p>
				</div>
			</CardContent>
			<CardFooter className="flex flex-col gap-2">
				<Button
					className="w-full bg-gradient-to-r from-purple-600 to-indigo-600 hover:from-purple-700 hover:to-indigo-700 text-white cursor-pointer"
					onClick={() => handlePayment("stripe")}
				>
					Pay with Stripe{" "}
					<BsStripe
						style={{
							width: "1.5em",
							height: "1.5em",
						}}
					/>
				</Button>
				<Button
					className="w-full bg-gradient-to-r from-blue-500 to-blue-600 hover:from-blue-600 hover:to-blue-700 text-white cursor-pointer"
					onClick={() => handlePayment("paypal")}
				>
					Pay with PayPal{" "}
					<BsPaypal
						style={{
							width: "1.5em",
							height: "1.5em",
						}}
					/>
				</Button>
				<Button
					className="w-full bg-gradient-to-r from-teal-400 to-teal-500 hover:from-teal-500 hover:to-teal-600 text-white cursor-pointer"
					onClick={() => handlePayment("mercadopago")}
				>
					Pay with Mercado Pago
				</Button>
			</CardFooter>
		</Card>
	);
}
