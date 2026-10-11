export default function debounce<T extends (...args: unknown[]) => void>(
	callback: T,
	ms = 200,
) {
	let id: ReturnType<typeof setTimeout>;
	const cancelDebounce = () => clearTimeout(id);

	function fn(...args: Parameters<T>) {
		cancelDebounce();
		id = setTimeout(() => callback(...args), ms);
	}

	fn.cancelDebounce = cancelDebounce;

	return fn;
}
