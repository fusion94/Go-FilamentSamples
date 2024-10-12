import unittest
from unittest.mock import patch, mock_open, call
import subprocess
from openscad_script import get_openscad_path, gen_samples

class TestOpenScadScript(unittest.TestCase):

    @patch('platform.system')
    def test_get_openscad_path_windows(self, mock_system):
        mock_system.return_value = 'Windows'
        self.assertEqual(get_openscad_path(), r'C:\Program Files\OpenSCAD\openscad.exe')

    @patch('platform.system')
    def test_get_openscad_path_linux(self, mock_system):
        mock_system.return_value = 'Linux'
        self.assertEqual(get_openscad_path(), 'openscad')

    @patch('os.path.exists', return_value=True)
    @patch('platform.system')
    def test_get_openscad_path_mac(self, mock_system, mock_exists):
        mock_system.return_value = 'Darwin'
        self.assertEqual(get_openscad_path(), '/Applications/OpenSCAD.app/Contents/MacOS/OpenSCAD')

    @patch('subprocess.run')
    @patch('os.path.exists', side_effect=[False, False])
    @patch('platform.system')
    def test_get_openscad_path_mac_with_command(self, mock_system, mock_exists, mock_run):
        mock_system.return_value = 'Darwin'
        mock_run.return_value = subprocess.CompletedProcess(args=['openscad', '-v'], returncode=0)
        self.assertEqual(get_openscad_path(), 'openscad')

    @patch('subprocess.run', side_effect=subprocess.CalledProcessError(1, ['openscad']))
    @patch('os.path.exists', side_effect=[False, False])
    @patch('platform.system')
    def test_get_openscad_path_mac_command_not_found(self, mock_system, mock_exists, mock_run):
        mock_system.return_value = 'Darwin'
        with self.assertRaises(SystemExit):
            get_openscad_path()

    @patch('builtins.open', new_callable=mock_open, read_data="Prusa,PLA,Red,210,60\n")
    @patch('os.makedirs')
    @patch('subprocess.run')  # Mock subprocess.run to prevent actual execution
    @patch('os.path.dirname', return_value='/dummy/path')
    def test_gen_samples(self, mock_dirname, mock_run, mock_makedirs, mock_open_file):
        # Mock the csv file content and file path
        test_csv = '/dummy/path/samples.csv'
        test_openscad = 'openscad'
        
        # Call the function with mock data
        gen_samples(test_csv, test_openscad)
        
        # Check if directories were created
        mock_makedirs.assert_called_once_with('/dummy/path/stl', exist_ok=True)
        
        # Check if subprocess.run was called with correct args
        expected_call = call([
            'openscad', '-o', '/dummy/path/stl/Prusa_PLA_Red_210_60.stl',
            '-D', 'BRAND="Prusa"', '-D', 'TYPE="PLA"', '-D', 'COLOR="Red"',
            '-D', 'TEMP_HOTEND="210"', '-D', 'TEMP_BED="60"',
            '/dummy/path/FilamentSamples.scad'
        ], check=True)
        
        mock_run.assert_has_calls([expected_call])

if __name__ == '__main__':
    unittest.main()
