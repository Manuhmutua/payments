FROM alpine
ADD payments /payments
ENTRYPOINT [ "/payments" ]
