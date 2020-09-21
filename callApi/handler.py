from flask import Flask
import json

app = Flask(__name__)

api_end_points = ['get_user_role', 'new_user', 'check_eligibilty', 'logout', 'forgot_password', 'find_educational_institution',
    'subscriber_dashboard', 'paid_subscriber', 'all_notifications', 'read_notifications', 'email_verification', 'upload_document',
    'id_upload', 'basic_lifestyle', 'religious_background', 'contact_details', 'family_details', 'education_career', 'about_me',
    'partner_basic_lifestyle', 'partner_religion_country', 'partner_education_career', 'choiceof_partner', 'country', 'partner_country',
    'country_codes', 'state', 'states_acc_country', 'cities', 'all_locations', 'caste', 'caste_group_value', 'general_settings',
    'deactivate_profile', 'contact_us', 'my_profile', 'other_profile', 'my_friends', 'add_friend', 'all_friends_category',
    'cancel_friend_request', 'deny_friend', 'accept_request', 'remove_friend', 'shortlisted_friend', 'unshortlisted_friend',
    'folder_listing', 'custom_folder', 'add_folder', 'rename_folder', 'albums', 'delete_album', 'change_permission',
    'update_profile_picture', 'subscription', 'media_coverages', 'success_stories', 'search_by_id', 'search_by_keyword',
    'advance_search', 'get_numbers', 'get_all_chats', 'generate_payment_hash_new', 'chat_details_new', 'send_chat_new',
    'update_status_chat', 'payment_bank_detail_send', 'send_otp_again', 'otp_verification', 'send_vertification', 'accept_again',
    'viewed_contacts', 'read_all_notifications', 'checkusername', 'get_updated_name', 'checkmobile']

@app.route('/<path>')
def api(path):
    
    if path in api_end_points:
        filename = 'jsonFiles/'+ path + ".json"
        with open(filename) as file:
            data = file.read()
            return data
    else:
        return None
