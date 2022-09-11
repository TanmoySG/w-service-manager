import nodemailer from "nodemailer"

export default class Mailer {
    constructor(configurations) {
        this.Server = configurations.server
        this.Port = configurations.port
        this.Sender = configurations.sender
        this.Password = configurations.password
    }


    async mail(to, subject, body, attachments) {
        let transporter = nodemailer.createTransport({
            host: this.Server,
            port: this.Port,
            secure: true, // true for 465, false for other ports
            auth: {
                user: this.Sender,
                pass: this.Password,
            },
        });

        // send mail with defined transport object
        let info = await transporter.sendMail({
            from: `${this.Sender}`, // sender address
            to: to.join(" , "), // list of receivers
            subject: subject, // Subject line
            html: body, // html body
            attachments: attachments
        });

        console.log("Message sent: %s", info.messageId);
    }

}
