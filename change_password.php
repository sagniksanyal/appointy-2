<?php
  require_once("../../../private/initialize.php");
  $page_title = "Staff - Password";
  require_login_admin();
  if(is_get_request()){
    $id = $_GET['id'];
    $result = find_staff_by_sid($id);
    if(!$result){
      redirect_to(url_for('/admin/staff/staff.php'));
    }
  }else if(is_post_request()){
    $staff['sid'] = $_POST['sid'];
    $staff['password'] = $_POST['password'];
    $result = update_staff_password($staff);
    if($result){
      redirect_to(url_for('/admin/staff/staff.php'));
    }
  }else{
    redirect_to(url_for('/admin/staff/staff.php'));
  }
  require_once(SHARED_PATH."/admin_header.php");
 ?>
 <link rel="stylesheet" href="<?php echo url_for("/stylesheets/password.css") ?>">
 <div class="wrap">
   <p>Update password for Staff with SID:<?php echo $id." "; ?>?</p>
   <table>
     <tr>
       <td>
         <strong>Name </strong>
       </td>
       <td>:&nbsp;
         <?php echo $result['name']; ?>
       </td>
     </tr>
     <tr>
       <td>
         <strong>Email </strong>
       </td>
       <td>:&nbsp;
         <?php echo $result['email']; ?>
       </td>
     </tr>
     <tr>
       <td>
         <strong>Mobile </strong>
       </td>
       <td>:&nbsp;
         <?php echo $result['phone']; ?>
       </td>
     </tr>
   </table>
   <form action="change_password.php" method="post">
     <input type="password" name="password" value="">
     <input type="hidden" name="sid" value="<?php echo $id; ?>">
     <input type="Submit" name="" value="Change Password">
   </form>
 </div>
 <?php
  require_once(SHARED_PATH."/footer.php");
  ?>
