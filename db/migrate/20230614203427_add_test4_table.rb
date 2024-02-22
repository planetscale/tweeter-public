class AddTest4Table < ActiveRecord::Migration[7.0]
  def change
    create_table :test4 do |t|
      t.string :name
      t.string :email
      t.timestamps
    end
  end
end
