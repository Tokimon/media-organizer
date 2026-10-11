type EventName = keyof DocumentEventMap;

type EventHandler<E extends EventName = EventName> = (
	this: EventSource | EventTarget,
	event: DocumentEventMap[E],
) => void;

export function on<E extends EventName>(
	context: typeof window | HTMLElement,
	eventName: E,
	handler: EventHandler<E>,
	options?: AddEventListenerOptions,
) {
	context.addEventListener(eventName, handler as EventListener, options);
	return () =>
		context.removeEventListener(eventName, handler as EventListener, options);
}
