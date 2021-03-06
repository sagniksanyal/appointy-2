<?php
require_once('../../private/initialize.php');
$page_title='Login-Admin';
// echo "<script>alert('Login Successful');</script>";
require_once(SHARED_PATH.'/header.php');
if(isset($_POST['aid']))
	{
    // echo "<script>alert('Login Successful');</script>";
		$aid=$_POST['aid'];
		$password=$_POST['password'];
		if($aid !='' && $password!='')
		{
			$admin=find_admin_by_aid($aid);
			$pass=$admin['pass'];
			if($pass!='')
			{
				if(password_verify($password,$pass))
				{
					login_admin($aid);
          // echo "<script>alert('Login Successful');</script>";
					redirect_to(url_for('/admin/admin_index.php'));
				}
				else
				{
					echo "<script>alert('Invalid Username or Password');</script>";
				}
			}
		}else{
			echo "<script>alert('Invalid Username or Password');</script>";
		}
	}
?>
<link rel="stylesheet" href="<?php echo url_for('/stylesheets/login.css');?>">
  <div class="wrap">
        <img class="f_image" src="<?php echo url_for('/images/admin.svg');?>" alt="logo">
        <form action="login.php" method="post" onsubmit="return validateForm()">
          <input class='login-form' onkeyup="nameCheck(this)" type="text" name="aid" id="username" value="" placeholder="Username">
          <span id="username_error" class="error hide"></span>
          <input class='login-form' onkeyup="passwordCheck(this)" type="password" name="password" id="password" value="password" placeholder="Password">
          <span id="password_error" class="error hide"></span>
          <input class="btn" type="submit" value="Submit"></input>
        </form>
  </div>
  <script src="<?php echo url_for('script/login.js')?>" defer>

  </script>
<?php include(SHARED_PATH.'/footer.php') ?>
