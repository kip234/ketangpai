package Handlers

const subject = "(0_0)用户中心"
const emailTimeLimit=600
const serviceTimeLimit=6
var body = `
		<html>
		<body>
   		<p>%+v</p>
   		<p>如果您并未发过此请求，则可能是因为其他用户误输入了您的电子邮件地址而使您收到这封邮件，那么您可以放心的忽略此邮件，无需进一步采取任何操作。</p>
   		<br>
   		<p>(0_0)用户中心</p>
   		<p></p><hr> <p></p>
   		<p> 除非您是收件人或负责向收件人转呈的人，本电子邮件及其附件严禁散发和复制。如果您误收此电子邮件，请删除此电子邮件及其附件。</p>
		</body>
		</html>
		`
// emmm... 将就一下吧