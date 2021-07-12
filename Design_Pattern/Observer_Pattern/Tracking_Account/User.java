public class User {
    private String email;
    private String ip;
    private LoginStatus status;

    public void setEmail(String email){
        this.email = email;
    }

    public void setIp(String ip){
        this.ip = ip;
    }

    public void setStatus(LoginStatus status){
        this.status = status;
    }

	public LoginStatus getStatus() {
		return this.status;
	}

	public String getEmail() {
		return this.email;
	}

	public String getIp() {
		return this.ip;
	}
}